# bisky

bisky compiles and runs your code for the appropriate language then judges it against the given test case.

it is broken up into two lambdas: the compiler lambda and the validator lambda.

the compiler lambda compiles and runs arbitrary code for a language you send to it and gives you the stdout/stderr

it only judges a single test case because we can just queue up compile jobs (even though this is slightly slower) it's more convenient until we separate "run all test cases" from "run single test cases"

# building/testing

Run `build.sh` then with the image id `docker run -p 9000:8080 0d02dbe8bf24`

# example request

```bash
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"Code":"print(517+1)", "Language": "Python"}' | jq
```
```bash
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"EncodedCode":"CmNsYXNzIFNvbHV0aW9uOgogICAgZGVmIHR3b1N1bShzZWxmLCBudW1zLCB0YXJnZXQpOgogICAgICAgIGQgPSB7fQogICAgICAgIGZvciBpIGluIHJhbmdlKDAsIGxlbihudW1zKSk6CiAgICAgICAgICAgIG4gPSBudW1zW2ldCiAgICAgICAgICAgIGsgPSB0YXJnZXQtbgogICAgICAgICAgICBpZiBrIGluIGQ6CiAgICAgICAgICAgICAgICByZXR1cm4gW2ksIGRba11dCiAgICAgICAgICAgIGRbbl0gPSBpCiAgICAgICAgcmV0dXJuIFtdCgk=", "Language": "Python", "Problem" :"two_sum", "TestCases": [{"Input": "[1,2,3,4,5]\n5", "ExpectedOutput": ["[0,1]"]},{"Input": "[1,2]\n2", "ExpectedOutput": ["[0,1]"]}]}' | jq
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
