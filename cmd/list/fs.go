package list

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	baseClass "github.com/Cyber-SiKu/cli/cmd/base"
	mds "github.com/Cyber-SiKu/cli/proto/curvefs/mds"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)
type (
	MdsAddr struct{
		Kind string		`mapstructure:"kind"`
		Curvefs string	`mapstructure:"mds.addr"`
		S3 string	`mapstructure:"s3"`
	}
)
type Fs struct {
	baseClass.CmdValue
}

var _ baseClass.CmdFunctions = (*Fs) (nil) //  检查 Fs 是否实现接口 CmdFunctions

func init() {
	
}

func NewFsCmd() *cobra.Command {
	fsCmd := &Fs{ baseClass.CmdValue{
		Use:	"fs",
		Short:	"list fs",
		Long:	"list fs",
	}}
	fsCmd.Cmd = baseClass.NewCmd(&fsCmd.CmdValue, fsCmd)
	return fsCmd.Cmd
}

func (c *Fs) Init(cmd *cobra.Command, args []string) error {
	c.Format = viper.GetString("format")
	return nil
}

func (c *Fs) RunCommand(cmd *cobra.Command, args []string) error {
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

	client := mds.NewMdsServiceClient(conn)
	request := &mds.ListClusterFsInfoRequest{}
	response, err := client.ListClusterFsInfo(context.Background(), request)
	if err != nil {
		fmt.Println(err)
	}
	m := protojson.MarshalOptions{
		Indent:          "  ",
		EmitUnpopulated: true,
	}
	c.Output, _ = m.Marshal(response)

	return nil
}

func (c *Fs) Print(cmd *cobra.Command, args []string) error {
	// fmt.Println(c.Format)
	if c.Format == "json" {
		fmt.Println(string(c.Output))
	} else {
		m := map[string]interface{}{}
		// Parsing/Unmarshalling JSON encoding/json
		json.Unmarshal(c.Output, &m)
		parseMap(m, 0)
	}
	return nil
}

func parseMap(aMap map[string]interface{}, level int) {
	for key, val := range aMap {
		for i := 0; i < level; i++ {
			fmt.Print("  ")
		}
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(val.(map[string]interface{}), level + 1)
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}), level + 1)
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}, level int) {
	for i, val := range anArray {
		for i := 0; i < level; i++ {
			fmt.Print("  ")
		}
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}), level+1)
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}), level+1)
		default:
			fmt.Println("Index", i, ":", concreteVal)
		}
	}
}
