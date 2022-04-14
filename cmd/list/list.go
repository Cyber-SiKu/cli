package list

import (
	"context"
	"fmt"
	"log"

	baseClass "github.com/Cyber-SiKu/cli/cmd/base"
	topology "github.com/Cyber-SiKu/cli/proto/curvefs/topology"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

type List struct {
	baseClass.CmdValue
}

func init() {
}

type Mds struct {
	Online []string `json:"online"`
	Offline []string	`json:"offline"`
	Leader []string		`json:"leader"`
}

func NewListCmd() *cobra.Command {
	listCmd := &List{ baseClass.CmdValue{
		Use:	"list",
		Short:	"list command",
		Long:	"list command",
	}}
	listCmd.Cmd = baseClass.NewCmd(&listCmd.CmdValue, listCmd)
	return listCmd.Cmd
}

func (l *List) Init(cmd *cobra.Command, args []string) error {
	l.AddSubCommand()
	return nil
}

func (l *List) RunCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		_, err := l.Cmd.ExecuteC()
		return err
	}
	mdsAddr := &MdsAddr{}
	err := viper.UnmarshalKey("curvefs", &mdsAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("mdsaddr: ", mdsAddr)
	fmt.Println("mds.addr:", mdsAddr.Curvefs)
	conn, err := grpc.Dial(mdsAddr.Curvefs, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

	client := topology.NewTopologyServiceClient(conn)
	request := &topology.StatMetadataUsageRequest{}
	response, err := client.StatMetadataUsage(context.Background(), request)
	if err != nil {
		fmt.Println(err)
	}
	m := protojson.MarshalOptions{
		Indent:          "  ",
		EmitUnpopulated: true,
	}
	l.Output, _ = m.Marshal(response)
	return nil
}

func (l *List) Print(cmd *cobra.Command, args []string) error {
	fmt.Println(string(l.Output))
	return nil
}

func (l *List) AddSubCommand() {
	l.Cmd.AddCommand(NewFsCmd())
}
