/*
 * Copyright 2022 Han Xin, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fuse2grpc

import (
	"context"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/hanwen/go-fuse/v2/fuse"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chiyutianyi/grpcfuse/pb"
)

func (s *server) StatFs(ctx context.Context, req *pb.StatfsRequest) (*pb.StatfsResponse, error) {
	var (
		out    fuse.StatfsOut
		header fuse.InHeader
	)
	grpc_logrus.Extract(ctx).WithFields(log.Fields{
		"nodeId": req.Input.NodeId,
	}).Debug("StatFs")
	toFuseInHeader(req.Input, &header)

	st := s.fs.StatFs(ctx.Done(), &header, &out)
	if st == fuse.ENOSYS {
		return nil, status.Errorf(codes.Unimplemented, "method StatFS not implemented")
	}
	if st != fuse.OK {
		return &pb.StatfsResponse{Status: &pb.Status{Code: int32(st)}}, nil
	}
	return &pb.StatfsResponse{
		Blocks:  out.Blocks,
		Bfree:   out.Bfree,
		Bavail:  out.Bavail,
		Files:   out.Files,
		Ffree:   out.Ffree,
		Bsize:   out.Bsize,
		NameLen: out.NameLen,
		Frsize:  out.Frsize,
		Status:  &pb.Status{Code: 0},
	}, nil
}
