package category

import (
	"fmt"
	"github.com/wpf1118/toolbox/tools/db"
	"github.com/wpf1118/toolbox/tools/flag"
	"testing"
)

func Test_categoryServ_List(t *testing.T) {
	opts := flag.NewDefaultMysqlOpts()
	opts.Endpoint = "www.zzrs.xyz:13305"
	db.MysqlInit(opts)
	s := NewCategoryServ()

	list, err := s.AllList()
	fmt.Println(err, len(list))

	treeList := s.TreeList(list, 0)

	fmt.Println(treeList)
}
