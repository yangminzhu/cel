# CEL playground

A command-line help tool for Common Expression Language [(CEL)](https://github.com/google/cel-spec).

## Examples

1. Evaluate a CEL expression using built-in sample attributes context
    ```bash
    $ make build && `make out` 'headers[":path"].startsWith("/info") && headers[":method"] in ["GET", "HEAD"] && (("x-id" in headers) ? (headers["x-id"]=="123456") : false)'
    protoc --go_out=. attributes/*.proto
    go build -o /home/ymzhu/go/out/yangminzhu/cel ./main
    2019/04/28 00:08:44 using attributes:
    clusters:
    - cluster-1
    - cluster-2
    - cluster-3
    headers:
      :host: httpbin
      :method: GET
      :path: /info/v1
      authorization: bearer 123456
      x-id: "123456"
    port: 8080
    sni: www.httpbin.com
    tls: true
    weighet:
    - 10
    - 20
    - 30
    - 40
    
    2019/04/28 00:08:44 evaluated successfully: true (bool)
    true
    ```

1. Parse a CEL expression and generate the AST
    ```
    $ make build && `make output` -ast 'headers[":path"].startsWith("/info") && headers[":method"] in ["GET", "HEAD"] && (("x-id" in headers) ? (headers["x-id"]=="123456") : false)'
    
    protoc --go_out=. attributes/*.proto
    go build -o /home/ymzhu/go/out/yangminzhu/cel ./main
    callExpr:
      args:
      - callExpr:
          args:
          - callExpr:
              args:
              - constExpr:
                  stringValue: /info
                id: "5"
              function: startsWith
              target:
                callExpr:
                  args:
                  - id: "1"
                    identExpr:
                      name: headers
                  - constExpr:
                      stringValue: :path
                    id: "3"
                  function: _[_]
                id: "2"
            id: "4"
          - callExpr:
              args:
              - callExpr:
                  args:
                  - id: "6"
                    identExpr:
                      name: headers
                  - constExpr:
                      stringValue: :method
                    id: "8"
                  function: _[_]
                id: "7"
              - id: "10"
                listExpr:
                  elements:
                  - constExpr:
                      stringValue: GET
                    id: "11"
                  - constExpr:
                      stringValue: HEAD
                    id: "12"
              function: '@in'
            id: "9"
          function: _&&_
        id: "13"
      - callExpr:
          args:
          - callExpr:
              args:
              - constExpr:
                  stringValue: x-id
                id: "14"
              - id: "16"
                identExpr:
                  name: headers
              function: '@in'
            id: "15"
          - callExpr:
              args:
              - callExpr:
                  args:
                  - id: "18"
                    identExpr:
                      name: headers
                  - constExpr:
                      stringValue: x-id
                    id: "20"
                  function: _[_]
                id: "19"
              - constExpr:
                  stringValue: "123456"
                id: "22"
              function: _==_
            id: "21"
          - constExpr:
              boolValue: false
            id: "23"
          function: _?_:_
        id: "17"
      function: _&&_
    id: "24"
    ```
