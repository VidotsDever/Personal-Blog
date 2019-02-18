export function convquery(str) {
    let obj = {};
    let parts = str.split("&");
    for (let index in parts) {
        let kws = parts[index].split("=");
        if(kws.length === 2) {
            obj[kws[0]] = kws[1];
        }
    }
    return obj
}