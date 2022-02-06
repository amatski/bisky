package codegen

const (
	CppTemplate = `
#include <iostream>
#include <vector>
#include <map>
#include <string>

using namespace std;

template < class T >
std::ostream& operator << (std::ostream& os, const std::vector<T>& v) 
{
    os << "[";
    for (typename std::vector<T>::const_iterator ii = v.begin(); ii != v.end(); ++ii)
    {
        os << *ii;
        if (ii+1 != v.end()) {
            os << ",";
        }
    }
    os << "]";
    return os;
}

${solution}

int main() {
    ${tests}
    return 0;
}
	`
	GoTemplate = `
package main

import "fmt"

${solution}

func main() {
	${tests}
}
	`
	JsTemplate = `
${solution}
${tests}
	`
	PyTemplate = `
${solution}
${tests}
	`
)

var (
	LanguageToTemplate = map[string]string{
		Python:     PyTemplate,
		Cpp:        CppTemplate,
		Javascript: JsTemplate,
		Go:         GoTemplate,
	}
)
