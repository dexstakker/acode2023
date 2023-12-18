# ADVENT CALENDAR OF CODE 2023
fobj = open(r"trebuchet.txt", "r")
content = fobj.readlines()
numlist = ['zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']

# Part A
total = 0
for w in content:
    w = w.strip()
    # # Part B needs to replace text of digits with the digits BEFORE processing the line
    # w = w.replace('one', '1')
    # w = w.replace('two', '2')
    # w = w.replace('three', '3')
    # w = w.replace('four', '4')
    # w = w.replace('five', '5')
    # w = w.replace('six', '6')
    # w = w.replace('seven', '7')
    # w = w.replace('eight', '8')
    # w = w.replace('nine', '9')

    print("Before Val = ", w)
    l = 0
    r = len(w)-1
    done = False
    while not done:
        if '0' <= w[l] <= '9':
            c1 = int(w[l])
            break
        nw = w[l:]
        for i in range(len(numlist)):
            if nw.startswith(numlist[i]):
                #nw.replace(numlist[i], str(i))
                c1 = i
                done = True
                break
        l = l + 1

    done = False
    while not done:
        if '0' <= w[r] <= '9':
            c2 = int(w[r])
            break
        nw = w[:r+1]
        for i in range(len(numlist)):
            if nw.endswith(numlist[i]):
                #nw.replace(numlist[i], str(i))
                c2 = i
                done = True
                break
        r = r - 1
    res = str(c1) + str(c2)
    total += int(res)
    print("Value = ", res)
print("TOTAL = ", total)