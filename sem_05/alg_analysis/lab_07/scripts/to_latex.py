fileName = '../data/param.txt'

with open(fileName, 'r') as f:
    a = f.readlines()

for s in a:
    r = ''
    while s != '':
        m = s[s.find('=') + 1:s.find(' ')]
        r += m + ' '
        if s == s[s.find(' ') + 1:]:
            r = r[:-1]
            s = ''
        s = s[s.find(' ') + 1:]
    # r = r.replace(' ', ' & ')
    # if r.split()[-1] == '0':
    #     print(r + ' \\\\')
    n = r.split()
    for i in range(len(n)):
        if i == (len(n) - 1):
            print(n[i], end='')
        else:
            print(n[i] + ',\t', end='')
    print()