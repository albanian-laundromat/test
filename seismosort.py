def issorted(l):
    k=len(l)-1;t=True
    for i in range(k):
        if l[i]>l[i+1]:t=False
    return t
def cycle(l):
    d = {};k = len(l);v = 0
    for i in range(k-1):
        if v not in d:
            d[v]=[]
        d[v].append(l[i])
        if l[i+1] > l[i]:
            v += 1
        elif l[i+1] < l[i]:
            v -= 1
    if v not in d:
        d[v]=[]
    d[v].append(l[k-1]);l = []
    for i in range(min(d.keys()),1+max(d.keys())):l += d[i]
    return l
def seismosort(l):
    while not issorted(l):
        print(l)
        l = cycle(l)
    print(l)
    return l
