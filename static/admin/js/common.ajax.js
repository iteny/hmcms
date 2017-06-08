/**
 * Created by iteny on 2016/5/15.
 */
$(function(){
    //排序
    $('button.ajax-sort').on('click', function (e) {
        e.preventDefault();
        var btn = $(this),
            form = btn.parents('form[name=commonSort]');
        if(btn.attr("disabledSubmit")){
            btn.text('请勿重复提交...').prop('disabled', true).addClass('disabled');
            return false;
        }
        btn.attr("disabledSubmit",true);
        var param = form.serialize();
        $.ajax({
            url: form.attr('action'),
            dataType:'json',
            type:'POST',
            data:param,
            beforeSend: function(){
                myload = layer.load(0,{time:3*1000});
            },
            success: function(data){
                layer.close(layer.load(1));
                if(!data.status){
                    admin.alert('提示信息',data.info,2,'3000');
                } else {
                    admin.countdown(3);
                    admin.alert('提示信息',data.info+'<div>程序将在<b style="color:red;" id="second_show">03秒</b>后为你跳转！</div>',1,'3000');
                    setTimeout(function(){
                        admin.reloadPage();
                    },3000);
                }
            },
            error: function(data){
                layer.close(layer.load(1));
                admin.alert('提示信息',data.responseText,1,'3000');
            }
        });
    });
    //单个删除
    $('.ajax-del').on('click', function (e) {
        e.preventDefault();
        var btn = $(this),
            title = btn.attr('data-title'),
            type = btn.attr('data-type'),
            href = btn.attr('href'),
            id = btn.attr('data-id');
        layer.confirm('是否删除'+type+'<span style="color:red;">['+title+']</span>', {icon: 3,offset:'200px', title:'删除'+type+'提示'}, function(index){
            layer.close(index);
            $.ajax({
                type: 'post',
                url: href,
                data : {id : id},
                dataType:"json",
                beforeSend: function(){
                    myload = layer.load(0,{time:3*1000});
                },
                success: function(data){
                    layer.close(layer.load(1));
                    if(data.status == true){
                        admin.countdown(3);
                        admin.alert('提示信息',data.info+'<div>程序将在<b style="color:red;" id="second_show">03秒</b>后为你跳转！</div>',1,'3000');
                        setTimeout(function(){
                            admin.reloadPage();
                        },3000);
                    } else {
                        admin.alert('提示信息',data.info,2,'3000');
                    }
                },
                error: function(data){
                    layer.close(layer.load(1));
                    admin.alert('提示信息',data.responseText,1,'3000');
                }
            });
        });
    });
    //批量删除
    $('button.ajax-batch-del').on('click', function (e) {
        e.preventDefault();
        if($("input:checkbox[name='ids[]']:checked").length <= 0){
            admin.alert('操作提示','至少选择一行数据进行操作',3,'5000');
            return false;
        }
        var btn = $(this),
            form = btn.parents('form[name=ids]'),
            type = btn.attr('data-type'),
            param = form.serialize();
            title = form.serializeArray();
            var ids=[];
            for(var i=0; i<title.length; i++){
                ids.push(title[i].value);
            }
        layer.confirm('是否删除ID为<span style="color:red;">['+ids+']</span>的'+type, {icon: 3,offset:'200px', title:'删除'+type+'提示'}, function(index){
            layer.close(index);

            $.ajax({
                url: btn.attr('href'),
                dataType:'json',
                type:'POST',
                data:param,
                beforeSend: function(){
                    myload = layer.load(0,{time:3*1000});
                },
                success: function(data){
                    layer.close(layer.load(1));
                    if(!data.status){
                        admin.alert('提示信息',data.info,2,'3000');
                    } else {
                        admin.countdown(3);
                        admin.alert('提示信息',data.info+'<div>程序将在<b style="color:red;" id="second_show">03秒</b>后为你跳转！</div>',1,'3000');
                        setTimeout(function(){
                            admin.reloadPage();
                        },3000);
                    }
                },
                error: function(data){
                    layer.close(layer.load(1));
                    admin.alert('提示信息',data.responseText,1,'3000');
                }
            });
        });
    });
    //全选的实现
    $(".check-all").click(function(){
        $(".ids").prop("checked", this.checked);
    });
    //悬停提示
    $('.bsn').poshytip({
        className: 'tip-darkgray',
        alignTo: 'target',
        alignX: 'right',
        alignY: 'center',
        offsetX: 5,
        showTimeout: 1100
    });
});
