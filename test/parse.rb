require 'json'

str=""
File.open("db/dbsetting.json", 'r') do |file|
	str = JSON.load(file.read())
end
puts ("#{str['DBname']}")
puts("#{str['User']}")
puts("#{str['Passphrase']}")
