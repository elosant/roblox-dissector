// +build nfqueue

package main

import (
	"fmt"
	"context"
	"syscall"

	"github.com/chifflier/nfqueue-go/nfqueue"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func PacketCallback(payload *nfqueue.Payload) int {
	fmt.Println("callback")
	packet := gopacket.NewPacket(payload.Data, layers.LayerTypeIPv4, gopacket.Default)

	fmt.Println("%s", packet.Dump())

	payload.SetVerdict(nfqueue.NF_ACCEPT)
	return 0
}

func CaptureFromNFQueue(ctx context.Context, session *CaptureSession) error {
	queue := new(nfqueue.Queue)
	queue.SetCallback(PacketCallback)
	queue.Init()

	queue.Unbind(syscall.AF_PACKET)
	queue.Bind(syscall.AF_PACKET)

	queue.CreateQueue(0)

	queue.Loop()

	// call queue.StopLoop() in another thread in response to a gui/menu action

	queue.DestroyQueue()
	queue.Close()

	return nil
}
