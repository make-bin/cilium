# deepequal-gen
A tool for auto-generating DeepEqual functions

## Overview

This tool is based on the Kubernetes® deepcopy-gen tool found here:
  https://github.com/kubernetes/gengo/tree/master/examples/deepcopy-gen

Given a list of input directories, it will generate DeepEqual methods that 
efficiently perform a full deep equal operation of each type. If these methods 
already exist (i.e., are predefined by the developer), they are used instead of
generating new ones.

The DeepEqual methods generated by this tool are presented as alternatives
to the reflect.DeepEqual which performs slower and requires that arrays/slices 
be sorted in the same order to report equality.  

The methods generated by this tool can override the default array/slice 
comparison approach to allow two slices with different element orders to 
report equality without first needing to sort the slices being compared.  This 
functionality is disabled by default. To enable this optional behaviour the type
must be annotated with the 'deepequal:unordered-array' tag.  For example, 
annotating the following type with this tag will generate a DeepEqual method 
that will return true regardless of the element order of the slices being 
compared as long as they both contain the equivalent elements.

```go
// deepequal-gen:unordered-array=true
type MyList []string
```

The methods generated by this tool can also override the default struct 
comparison approach to provide custom behaviour.  For applications that have a
need to express certain fields as optional and want to ignore the result of 
comparing those specific fields if the left hand operand is a nil pointer then 
the 'deepequal-gen:ignore-nil-fields' tag can be applied to the struct 
definition.  For example, given the following struct it may be desirable to
consider to instances as equal regardless of the value of the dereferenced 'age'
field.  

```go
// +k8s:deepequal-gen:ignore-nil-fields=true
type MyStruct struct {
    Name string `json:"name"`
    Address string `json:"address"`
    Age *int `json:"age,omitempty"`
}

a := MyStruct{Name: "John", Address: "Somewhere"}
ageB := 30
b := MyStruct{Name: "John", Address: "Somewhere", Age: &ageB}
ageC := 31
c := MyStruct{Name: "John", Address: "Somewhere", Age: &ageC}

a.DeepEqual(&b) == true
b.DeepEqual(&c) == false
```
 
All generation is governed by comment tags in the source.  Any package may
request DeepEqual generation by including a comment in the file-comments of
a doc.go file, of the form:
  deepequal-gen=package

DeepEqual functions can be generated for individual types, rather than the
entire package by specifying a comment on the type definition of the form:
  deepequal-gen=true

When generating for a whole package, individual types may opt out of
DeepEqual generation by specifying a comment on the type definition of the
form:
  deepequal-gen=false

**Warning:**  This module should be considered experimental.  It was developed and
tested with a specific set of usecases in mind.  It should not be considered
a complete implementation that will handle all possible type implementations. If
you plan on using this module you should implement a comprehensive set of unit
tests in your application to ensure that the behaviour that your application
requires is implemented as expected.


## Project License

The license for this project is the Apache 2.0 license. Text of the Apache 2.0
license and other applicable license notices can be found in the LICENSE file
in the top level directory. Each source file should include a license notice
that designates the licensing terms for the respective file.


## Legal Notices

All product names, logos, and brands are property of their respective owners.
All company, product and service names used in this software are for
identification purposes only. Wind River is a registered trademark of Wind River
Systems, Inc.  Kubernetes is a registered trademark of Google Inc.

Disclaimer of Warranty / No Support: Wind River does not provide support and
maintenance services for this software, under Wind River’s standard Software
Support and Maintenance Agreement or otherwise. Unless required by applicable
law, Wind River provides the software (and each contributor provides its
contribution) on an “AS IS” BASIS, WITHOUT WARRANTIES OF ANY KIND, either
express or implied, including, without limitation, any warranties of TITLE,
NONINFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE. You are
solely responsible for determining the appropriateness of using or
redistributing the software and assume any risks associated with your exercise
of permissions under the license.