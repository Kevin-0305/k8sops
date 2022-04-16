export function dictToList (obj) {
    let list 
    list = []
    for (var i in obj){
        list.push({"key":i,"value":obj[i]})
    }
    return list
}

export function listToDict (arr) {
    let dict
    dict = {}
    for (const i of arr){
        console.log(i)
        dict[i.key]=i.value
    }
    return dict
}