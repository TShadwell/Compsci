fibb = lambda n: (n if n >=0 else fibb(n+2)-fibb(n+1))if n <2 else fibb(n-1)+fibb(n-2)
