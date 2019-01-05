require 'json'
json_str=""
File.open("httpdata.json", 'r') do |file|
	json_str = JSON.load(file.read())
end

INDEX_STATUS_CODE=1
INDEX_BODY=0
dotest=0
failed=0
mapped = json_str.map do |testtitle, testdata|
	dotest=dotest+1
	if testdata.has_key?('body_file')
		body_str="-d @#{testdata['body_file']} -H \"Content-Type: application/json\""
	else
		body_str=""
	end
	puts("Test:#{testtitle}")
	result=`curl http://localhost:3300#{testdata['uri']} -w 'http_code:%{http_code}\n' -X #{testdata['method']} #{body_str} 2> /dev/null`.split(/http_code:/)
	#result 0=>body, 1=>status code
	if result[INDEX_STATUS_CODE].to_i != testdata['result_status'].to_i
		puts("Failed to check status code, expected_value:#{testdata['result_status']}, result:#{result[INDEX_STATUS_CODE]}")
		failed=failed+1
	end

	if result[INDEX_STATUS_CODE].to_s == "200"
		if !testdata.has_key?('result_body')
			puts("Failed to check result body")
			failed=failed+1
		else
			expected_value = JSON.load(testdata['body_file'])
			result_value = JSON.load(result[INDEX_BODY])
			if expected_value != result_value
				puts("Failed to check result body, ")
				puts(expected_value)
				puts(result_value)
				failed=failed+1
			end
		end
	end
end
puts("Test done:failed/all=#{failed}/#{dotest}")
