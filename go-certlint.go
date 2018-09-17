/* go-certlint - Go wrapper for github.com/awslabs/certlint
 * Written by Rob Stradling
 * Copyright (C) 2018 COMODO CA Limited
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package certlint

import (
	"unsafe"
)

/*
#cgo pkg-config: ruby-2.3

#include <ruby.h>

typedef struct ruby_shared_data {
	VALUE obj;
	ID method_id;
	int nargs;
	VALUE args[4];
} ruby_shared_data;

static VALUE ruby_callback(VALUE ptr)
{
	ruby_shared_data* data = (ruby_shared_data*)ptr;
	return rb_funcall2(data->obj, data->method_id, data->nargs, data->args);
}

VALUE rescue_require(VALUE data, VALUE err)
{
	fprintf(stderr, "Error: [[[%s]]]\n", RSTRING(rb_obj_as_string(err))->as.heap.ptr);
	return Qnil;
}

void init_certlint(void)
{
	ruby_init();
	ruby_script("go-certlint");
	ruby_init_loadpath();
	ruby_incpush("/usr/local");
	ruby_incpush("/usr/local/certlint/lib");
	ruby_incpush("/usr/local/certlint/ext");
	rb_require("enc/encdb");
	rb_require("enc/trans/transdb");
	rb_require("certlint");
	rb_eval_string("def runcablint(raw) return CertLint::CABLint.lint(raw).join(\"\n\") end");
	rb_eval_string("def runcertlint(raw) return CertLint.lint(raw).join(\"\n\") end");
}

char* lint(char* linter, unsigned char* cert_buffer, size_t cert_len)
{
	ruby_shared_data rbdata;
	rbdata.obj = 0;
	rbdata.method_id = rb_intern(linter);
	rbdata.nargs = 1;
	rbdata.args[0] = rb_str_new((char*)cert_buffer, cert_len);
	VALUE t_result = rb_rescue(&ruby_callback, (VALUE)&rbdata, &rescue_require, Qnil);
	return RSTRING_PTR(t_result);
}
*/
import "C"

func Init() {
	C.init_certlint()
}

func Cablint(cert_der []byte) string {
	return C.GoString(C.lint(C.CString("runcablint"), (*C.uchar)(unsafe.Pointer(&cert_der[0])), (C.ulong)(len(cert_der)))) + "\n"
}

func Certlint(cert_der []byte) string {
	return C.GoString(C.lint(C.CString("runcertlint"), (*C.uchar)(unsafe.Pointer(&cert_der[0])), (C.ulong)(len(cert_der)))) + "\n"
}

func Finish() {
	C.ruby_cleanup(0)
}
