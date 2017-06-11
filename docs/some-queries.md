{
  "code": "Success",
  "message": "Done",
  "uids": {
    "class": "0x9aac7ab7df2e4874",
    "x": "0xea4dfda5c29aef61",
    "y": "0xfc4eba7c1fd6f776"
  }
}


{
 class(id:0x9aac7ab7df2e4874) {
  name
  student {
   name
   planet
   friend {
    name
   }
  }
 }
}


mutation {
  set {
    <alice> <name> "Alice" .
    <lewis-carrol> <died> "1998" .
  }
}

query {
  me(id:alice) {
    name
  }
}


mutation {
  delete {
     <lewis-carrol> <died> "1998" .
  }
}
