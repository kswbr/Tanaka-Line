const PROTO_PATH = __dirname + '/proto/tanaka.proto';
const protoLoader = require('@grpc/proto-loader')
const grpc = require('grpc');

const packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
   }
)

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition)
const tanakaProto = protoDescriptor.Tanaka
function main() {
  const client = new tanakaProto('scraper:8080',
                                       grpc.credentials.createInsecure());
  client.GetNews({}, function(err, response) {
    if (err) {
      throw err
    }
    console.log(response)
  });
}
main()
