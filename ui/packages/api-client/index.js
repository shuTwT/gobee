import * as path from "node:path";
import * as process from "node:process";
import { generateApi } from "swagger-typescript-api";

await generateApi({
   input: path.resolve(process.cwd(),"..","..","..", "./swagger.json"),
   url: "http://localhost:13000/swagger/doc.json",
   output: path.resolve(process.cwd(), "./dist"),
   fileName: "index.ts",
    // silent: true,
    // modular:true,
    httpClientType: "axios",
    toJS:true,
    singleHttpClient: true,
    generateClient: true,
    generateRouteTypes: true,
    sortTypes: true,
    hooks:{

    },
    templates:path.resolve(process.cwd(), "./templates"),
   });
