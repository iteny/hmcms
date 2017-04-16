<style>
.hm_loginForm .el-form-item.is-required .el-form-item__label:before{
    content: '';    
    margin-right: 0px;
}
.hm_loginForm .title {
    margin: 0px auto 30px auto;
    text-align: center;
    color: #888;
}
</style>
<template>
  <el-form :model="hmForm" :rules="rules" ref="hmForm" label-width="100px" class="hm_loginForm">
    <h3 class="title">后台登录</h3>
    <el-form-item label="用户名" prop="username">
      <el-input v-model="hmForm.username"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input v-model="hmForm.password"></el-input>
    </el-form-item>
    
    <el-form-item>
      <el-button type="primary" @click="submitForm('hmForm')" style="width:100%;">立即创建</el-button>
      <!-- <el-button @click="resetForm('hmForm')">重置</el-button> -->
    </el-form-item>
  </el-form>
</template>
<script>
  export default {
    data() {
      return {
        hmForm: {
          username: '',
          password: '',
        },
        rules: {
          username: [
            { required: true, message: '请输入活动名称', trigger: 'blur' },
            { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
          ],
          password: [
            { required: true, message: '请输入活动名称', trigger: 'blur' },
            { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
          ],
          
        }
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            //发送get请求
            // var item = {username:this.hmForm.username,password:this.hmForm.password};
            // this.$http.post('/admin',item).then(function(res){
            //     alert(res.body);    
            // },function(){
            //     alert('请求失败处理');   //失败处理
            // });
             $.ajax({
                type: 'post',
                url: '/admin',
                data : {username:this.hmForm.username,password:this.hmForm.password},
                dataType:"json",
                beforeSend: function(){
                    // myload = layer.load(0,{time:3*1000});
                },
                // success: function(msg){
                //     if(msg.status === 1){
                //         admin.success(msg.info,'#loginsubmit');
                //         setTimeout(function(){window.location.href=redirect}, 3000);
                //     }else{
                //         layer.close(myload);
                //         admin.error(msg.info,'#loginsubmit');
                //         formLogin.attr('disabledSubmit','');
                //         // refreshs();
                //         $('.yanzheng_img').eq(0).click();
                //         $('#verify').val('')
                //     }                    
                // },
                // error: function(XMLHttpRequest, textStatus, errorThrown){
                //     admin.error('网络连接异常！','#loginsubmit');
                //     return false;
                // }
            });
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    }
  }
</script>