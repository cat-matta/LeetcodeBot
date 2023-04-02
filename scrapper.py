from bs4 import BeautifulSoup
import requests

request = requests.get('https://leetcode-company-tagged.vercel.app/hudson-river-trading/')
html_body = request.text

soup = BeautifulSoup(html_body, 'html.parser')
# print(soup.prettify())

table = soup.find('table')
elements = (table.find_all('tr'))

# print(elements)
out = []

table_row = {}
final = []
for i in range(1, len(elements)):
	item = elements[i]
	# print("QUESTION: ", i)
	row = item.find_all('td')
	# print(row, end='\n')
	out = []

	for thing in row:
		url = thing.find('a')
		if thing.text:
			# print(thing.text)
			out.append(thing.text)
		if url:
			# print(url['href'])
			out.append(url['href'])

	# print(out)
	table_row['Number'] = out[0]
	table_row['Title'] = out[1]
	table_row['Link'] = out[2]
	table_row['Difficulty'] = out[3]
	table_row['Frequency'] = out[4]
	final.append(table_row)
	# print("-"*10)

print(final)



