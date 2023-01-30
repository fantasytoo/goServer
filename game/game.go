/**
@author: Administrator
@Date: 2022-11-18-0018
@Note
**/
package game

import "github.com/lonng/nano"

func StartUp() {
	nano.Listen("127.0.0.1:10011")
}
