q = <<-'GRAPHQL'
mutation {
  set {
    <a> <name> "alice" .
    <b> <name> "bob" .
    <a> <age> "10" .
    <b> <age> "20" .
    <b> <boss_of> <a> .
    <a> <address> "London" .
    <b> <address> "San Francisco" .
  }
}
GRAPHQL


q = <<-'GRAPHQL'
{
# get questions to show in the home page
  var(func: gt(count(question.body), 0)) {
    ac as count(answer)
    score as math( ac * 0.2)
  }

  questions(id: var(score), orderasc: var(score), first: 20) {
    _uid_
    question.body
    question.title
    question.created_at
    question.written_by {
      user_name
    }
    answer_count: count(answer)
  }
}
GRAPHQL