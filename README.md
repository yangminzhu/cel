# cel

A command-line tool to evaluate CEL expression and also print the parsed AST.

## Example

```bash
$ go run main/main.go -ast  'a<=b && c==true && (d.contains("bye") || e[0]+e[1]>=e[2]) && ("f2" in f) && g["g2"].startsWith("g")' 
2019/04/27 01:20:48 generating AST for a<=b && c==true && (d.contains("bye") || e[0]+e[1]>=e[2]) && ("f2" in f) && g["g2"].startsWith("g")
callExpr:
  args:
  - callExpr:
      args:
      - callExpr:
          args:
          - callExpr:
              args:
              - id: "1"
                identExpr:
                  name: a
              - id: "3"
                identExpr:
                  name: b
              function: _<=_
            id: "2"
          - callExpr:
              args:
              - id: "4"
                identExpr:
                  name: c
              - constExpr:
                  boolValue: true
                id: "6"
              function: _==_
            id: "5"
          function: _&&_
        id: "7"
      - callExpr:
          args:
          - callExpr:
              args:
              - constExpr:
                  stringValue: bye
                id: "10"
              function: contains
              target:
                id: "8"
                identExpr:
                  name: d
            id: "9"
          - callExpr:
              args:
              - callExpr:
                  args:
                  - callExpr:
                      args:
                      - id: "11"
                        identExpr:
                          name: e
                      - constExpr:
                          int64Value: "0"
                        id: "13"
                      function: _[_]
                    id: "12"
                  - callExpr:
                      args:
                      - id: "15"
                        identExpr:
                          name: e
                      - constExpr:
                          int64Value: "1"
                        id: "17"
                      function: _[_]
                    id: "16"
                  function: _+_
                id: "14"
              - callExpr:
                  args:
                  - id: "19"
                    identExpr:
                      name: e
                  - constExpr:
                      int64Value: "2"
                    id: "21"
                  function: _[_]
                id: "20"
              function: _>=_
            id: "18"
          function: _||_
        id: "22"
      function: _&&_
    id: "23"
  - callExpr:
      args:
      - callExpr:
          args:
          - constExpr:
              stringValue: f2
            id: "24"
          - id: "26"
            identExpr:
              name: f
          function: '@in'
        id: "25"
      - callExpr:
          args:
          - constExpr:
              stringValue: g
            id: "32"
          function: startsWith
          target:
            callExpr:
              args:
              - id: "28"
                identExpr:
                  name: g
              - constExpr:
                  stringValue: g2
                id: "30"
              function: _[_]
            id: "29"
        id: "31"
      function: _&&_
    id: "33"
  function: _&&_
id: "27"


2019/04/27 01:20:48 using sample attributes: map[a:10 b:20 c:true d:hello e:[11 22 33] f:[f1 f2] g:map[g1:g11 g2:g22]]

2019/04/27 01:20:48 evaluated expression: a<=b && c==true && (d.contains("bye") || e[0]+e[1]>=e[2]) && ("f2" in f) && g["g2"].startsWith("g")
true
```
