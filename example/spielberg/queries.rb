class Snippet
  attr_accessor :code
  attr_accessor :name
end

class Registry
  attr_accessor :snippets

  def self.instance
    @instance ||= self.new
  end
  def self.add(name:, code:)
    self.instance.add(name, code)
  end

  def add(name, code)
    Snippet.new.tap do |s|
      s.name = name
      s.code = code
      snippets << s
    end
  end

  def inspect
    snippets.each do |s|
      puts "=== #{s.name}:"
      puts "#{s.code}"
    end
    nil
  end

  def snippets
    @snippets||= []
  end
end


r = Registry.new
r.add "moviesSchema", <<-'GRAPHQL'
mutation {
  schema {
    movie.name: string @index .
    person.name: string @index .
    movie.release_date: date @index .
    movie.revenue: float .
    movie.running_time: int .
  }
}
GRAPHQL

r.add "allMovies", <<-'GRAPHQL'
query {
	director(func:allofterms(name, "michael bay")) {
		name@en
		director.film (orderdesc: initial_release_date) {
			name@en
			initial_release_date
		}
	}
}
GRAPHQL

puts r.inspect
