import json
import pymongo

client = pymongo.MongoClient('localhost')
db = client['webCourse']
collection = db['film']

with open("./films_all.json", encoding="utf-8") as jf:
    str = jf.read()
    data = []
    data.extend(json.loads(str))
    collection.insert_many(data)

print(collection.count_documents({}))
client.close() #写完关闭连接
