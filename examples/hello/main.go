/* main.go - Example for goemacs API

Copyright (C) 2016 Yann Hodique <yann.hodique@gmail.com>.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or (at
your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.  */

package main

// int plugin_is_GPL_compatible;
import "C"

import "github.com/sigma/goemacs"

func init() {
	goemacs.Register(initModule)
}

func initModule(env *goemacs.Environment) {
	stdlib := env.StdLib()
	stdlib.Message("hello from go")
}

func main() {}
