let errMsg = {
    //检查特定字段是否必填
    required: {
        msg: '这个字段不能为空',
        test: function (obj, load) {
            return obj.value.length > 0 && obj.value != obj.defaultValue;
        }
    },
    //确保字段内容是正确的email地址
    email: {
        msg: 'email格式错误!',
        test: function (obj) {//确保有内容的输入并符合email地址的格式
            return !obj.value || /^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/.test(obj.value);
        }
    },
    //确保字段内容是电话号码并将其自动格式化
    phone: {
        msg: 'Not a valid phone number',
        test: function (obj) {
            var m = /(\d{3}).*(\d{3}).*(\d{4})/.exec(obj.value);
            if (m) {
                obj.value = "(" + m[1] + ")" + m[2] + "-" + m[3];
                return !obj.value || m;
            }
        }
    },
    //确保字段内容符合MM/DD/YYYY的时间格式
    date: {
        msg: 'Not a valid date.',
        test: function (obj) {
            return !obj.value || /^\d{2}\/d{2}\/d{2,4}$/.test(obj.value);
        }
    },
    url: {
        msg: 'Not a valid URL.',
        test: function (obj) {
            return !obj.value || /^(f|ht){1}(tp|tps):\/\/([\w-]+\.)+[\w-]+(\/[\w- ./?%&=]*)?/.test(obj.value);
        }
    },
    zifu: {
        msg: '5个字符~15个字符',
        test: function (obj) {
            return !obj.value || /^\S{5,15}$/.test(obj.value);
        }
    }

};
//显示表单内特定字段的错误信息
function showErrors(elem, errors,leng) {
    var next = elem.nextSibling;//获取当前字段的下一个元素
    if (next && (next.nodeName != 'div' && next.className != 'error')) {
        next = document.createElement('div');
        next.className = 'error';
        elem.parentNode.insertBefore(next, elem.nextSibling);
    }
    // next.className = 'errors';
    //有一个包含错误的容器引用，我们可以遍历所有的错误信息
    var li = document.createElement('span');
	// for (var i = 0; i < errors.length; i++) {
 //    //为每一个错误信息创建新的li包裹器
 //        li.innerHTML += errors[i] + '</br>';
        
 //    }
    if(next.firstChild == null){
    	next.appendChild(li);//并插入到dom中
    }
    // li = document.getElementsByTagName('span');
    
    console.info(errors);
    if(errors.length > 0 ){
    	next.firstChild.innerHTML = null;
        for (var i = 0; i < errors.length; i++) {
	    //为每一个错误信息创建新的li包裹器
	    	console.info(errors[i]);
	        next.firstChild.innerHTML += errors[i] + '</br>';
	        
	    }
	}
 //    if(errors.length > 0 ){
 //    	li.innerHTML = null;
	//     for (var i = 0; i < errors.length; i++) {
	//     //为每一个错误信息创建新的li包裹器
	//         li.innerHTML += errors[i] + '</br>';
	        
	//     }
	// }
}
//验证单个字段的内容
function validateField(elem, load) {
    var errors = [];
    for (var name in errMsg) {
        var re = new RegExp("(^|\\s)" + name + "(\\s|$)");   
        // console.info(name);     
        // console.info(errMsg[name]);
        if (re.test(elem.className) && !errMsg[name].test(elem, load)) {//如果没有通过验证，把错误信息增加到列表中
            errors.push(errMsg[name].msg);
        }

    }
    console.info(errors);
    if (errors.length) {//如果存在错误信息，则显示出来
        showErrors(elem, errors);
    }
    
    var nex = elem.nextSibling;
    if(errors.length == 0 && nex.className == 'error'){
    	nex.parentNode.removeChild(nex);
    }
    return errors.length > 0;
}
export default function validate(
	elem
){
	
    let me = validateField(elem) ? 'false' : 'true';
    return me;
}