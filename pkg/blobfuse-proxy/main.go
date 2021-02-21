/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	"google.golang.org/grpc"
	mount_azure_blob "sigs.k8s.io/blob-csi-driver/pkg/blobfuse-proxy/pb"
)

func init() {
	_ = flag.Set("logtostderr", "true")
}

var (
	blobfuseProxyEndpoint = flag.String("blobfuse-proxy-endpoint", "unix://tmp/blobfuse-proxy.sock", "blobfuse-proxy endpoint")
)

type MountServer struct {
	mount_azure_blob.UnimplementedMountServiceServer
}

// NewMountServer returns a new Mountserver
func NewMountServiceServer() *MountServer {
	return &MountServer{}
}

// MountAzureBlob mounts an azure blob container to given location
func (server *MountServer) MountAzureBlob(ctx context.Context,
	req *mount_azure_blob.MountAzureBlobRequest,
) (resp *mount_azure_blob.MountAzureBlobResponse, err error) {

	log.Printf("received request: Mounting the container %s to the path %s \n", req.GetContainerName(), req.GetTargetPath())
	resp = &mount_azure_blob.MountAzureBlobResponse{Err: ""}
	args := fmt.Sprintf("%s --tmp-path=%s --container-name=%s", req.GetTargetPath(), req.GetTmpPath(), req.GetContainerName())
	cmd := exec.Command("blobfuse", strings.Split(args, " ")...)
	// TODO: take mount args being passed from storage class

	cmd.Env = append(os.Environ(), "AZURE_STORAGE_ACCOUNT="+req.GetAccountName())
	cmd.Env = append(cmd.Env, "AZURE_STORAGE_ACCESS_KEY="+req.GetAccountKey())
	err = cmd.Run()
	if err != nil {
		log.Println("blobfuse mount failed")
		resp.Err = err.Error()
		return resp, err
	}
	log.Println("successfully mounted")
	return resp, nil
}

func runGRPCServer(
	mountServer mount_azure_blob.MountServiceServer,
	enableTLS bool,
	listener net.Listener,
) error {
	serverOptions := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOptions...)

	mount_azure_blob.RegisterMountServiceServer(grpcServer, mountServer)

	log.Printf("Start GRPC server at %s, TLS = %t", listener.Addr().String(), enableTLS)
	return grpcServer.Serve(listener)
}

func parseEndpoint(ep string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(ep), "unix://") || strings.HasPrefix(strings.ToLower(ep), "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("Invalid endpoint: %v", ep)
}

func main() {
	flag.Parse()
	proto, addr, err := parseEndpoint(*blobfuseProxyEndpoint)
	if err != nil {
		log.Fatal(err.Error())
	}

	if proto == "unix" {
		addr = "/" + addr
		if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
			log.Fatalf("Failed to remove %s, error: %s", addr, err.Error())
		}
	}

	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	mountServer := NewMountServiceServer()

	log.Printf("Listening for connections on address: %#v\n", listener.Addr())
	if err = runGRPCServer(mountServer, false, listener); err != nil {
		log.Fatalf("Listening for connections on address: %#v, error: %v", listener.Addr(), err)
	}
}
