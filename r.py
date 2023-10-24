s = input()
x = input()
s1 = s.lower()
x1 = x.lower()
ans = ''
while x1 in s1:
    s1 = s1.replace(x1, ' '*len(x1))
for i in range(len(s)):
    if int(ord(s1[i])) == int(ord(s[i])) or int(ord(s1[i]) - 32) == int(ord(s[i])) or int(ord(s1[i])) == int(ord(s[i]) - 32):
        ans += s[i]
print(ans)
