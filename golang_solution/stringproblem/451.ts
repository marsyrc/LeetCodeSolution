function frequencySort(s: string): string {
    const map = new Map();
    const arr = [];
    let res = "";
    for (let l of s) {
        if (map.has(l)) {
            map.set(l, map.get(l) + 1);
        } else {
            map.set(l, 1);
        }
    }

    for (let [key, value] of map) {
        if (!arr[value]) {
            arr[value] = [key];
        } else {
            arr[value].push(key);
        }
    }

    for (let i = arr.length - 1; i > 0; i--) {
        if (arr[i]) {
            for (let l of arr[i]) {
                res += l.repeat(i);
                res += l.repeat(i);
            }
        }
    }
    return res;
};
