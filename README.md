# bisky

bisky compiles and runs your code for the appropriate language then judges it against the given test case.

it is broken up into two lambdas: the compiler lambda and the validator lambda.

the compiler lambda compiles and runs arbitrary code for a language you send to it and gives you the stdout/stderr

it only judges a single test case because we can just queue up compile jobs (even though this is slightly slower) it's more convenient until we separate "run all test cases" from "run single test cases"

pass in the test case id and then it will look up the test case (stdin, stdout aka input, expected output) to do an equality check against 

if it is empty it will take in the test case directly in another field

yes but how will it take in input if they dont read from stdin, u have to call their code with the variables somehow

comparing the output is easy we just look it up from dynamodb but taking in input is significantly trickier
also what if they call stdin and halt the program?

you pretty much either need stdin hooks or to hook it via writing your own dumper but stdin is easier tbh
just make it impossible to override or guess the variables and always put it at the end of the file as a hook. then just write a test for each language against all the test cases with your hook and it should be verifiable. also the hooks can be generated automatically via code given just the variables and types

it's nice to do it this way too because then you can have your hooks do all the validation and return a code if it matches..or you could match on stdout directly (PER LANGUAGE?). matching on stdout is worse because when we match trees we want the actual in memory data structures

then all we have to print via stdout is Y or N and the code itself will know the result in stdout

first line via stdout will be their output then the second line is `Y` or `N` for passing or failing

for most languages you can just literally inject the data via code using sprintf, i believe its like that for every language except go and C?. give them super random scrambled names so redefinition isnt an issue

# how it works

When you generate code with an answer in any language we append a hook block for that specific language which hooks the code to generate a new program which prints the actual answer for a testcase in stdout then we use that output in the validator to ultimately judge the test case

roughly this:

```
compile(solution + genCodeWithValidator(hook(testcase_input, language))) -> their answer (stdout) -> validator(stdout, testcase_output)
```

This way we only write hooks for each language once and then the validator exactly once (and not multiple validators per language) without knowing the original answer's source language.

Also note that the hooks can also be programatically generated if the problem input and output format are POD types, so most problems can we just be parameterized by `testcase_input`, `testcase_output`, and `solution`.

# building/testing

Run `build.sh` then with the image id `docker run -p 9000:8080 0d02dbe8bf24`

# example request

```bash
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"Code":"print(517+1)", "Language": "Python"}' | jq
```
```bash
cub@mbp ~/go/src/amatski/bisky (master)$ curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"EncodedCode":"CmNsYXNzIFNvbHV0aW9uOgogICAgZGVmIHR3b1N1bShzZWxmLCBudW1zLCB0YXJnZXQpOgogICAgICAgIGQgPSB7fQogICAgICAgIGZvciBpIGluIHJhbmdlKDAsIGxlbihudW1zKSk6CiAgICAgICAgICAgIG4gPSBudW1zW2ldCiAgICAgICAgICAgIGsgPSB0YXJnZXQtbgogICAgICAgICAgICBpZiBrIGluIGQ6CiAgICAgICAgICAgICAgICByZXR1cm4gW2ksIGRba11dCiAgICAgICAgICAgIGRbbl0gPSBpCiAgICAgICAgcmV0dXJuIFtdCgk=", "Language": "Python", "Problem" :"two_sum", "TestCases": [{"Input": "[1,2,3,4,5]\n5", "ExpectedOutput": ["[0,1]"]},{"Input": "[1,2]\n2", "ExpectedOutput": ["[0,1]"]}]}' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   643  100   121  100   522   6368  27473 --:--:-- --:--:-- --:--:-- 35722
{
  "Results": [
    {
      "Stdout": "[2, 1]",
      "Passed": false,
      "Elapsed": 1000000000
    },
    {
      "Stdout": "[]",
      "Passed": false,
      "Elapsed": 1000000000
    }
  ]
}
```