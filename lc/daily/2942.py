class Solution:
	def findWordsContaining(self, words: List[str], x: str) -> List[int]:
		res = []
		i = 0
		for word in words:
			if x in word:
				res.append(i)
			i += 1
		return res
