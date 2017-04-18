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
      <el-button type="primary" @click="submitForm('hmForm')" style="width:100%;" v-loading.fullscreen.lock="fullscreenLoading">立即创建</el-button>
      <!-- <el-button @click="resetForm('hmForm')">重置</el-button> -->
    </el-form-item>
  </el-form>
</template>
<script>
  export default {
    data() {
      return {
        fullscreenLoading: false,
        hmForm: {
          username: '',
          password: '',
        },
        rules: {
          username: [
            { required: true, message: '请输入用户名', trigger: 'blur' },
            { min: 6, max: 15, message: '长度在 6 到 15 个字符', trigger: 'blur' }
          ],
          password: [
            { required: true, message: '请输入密码', trigger: 'blur' },
            { min: 8, max: 15, message: '长度在 8 到 15 个字符', trigger: 'blur' }
          ],
          
        }
      };
    },
    methods: {

      submitForm(formName) {
        
        this.$refs[formName].validate((valid) => {
          if (valid) {
            var e = this;
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
                    e.fullscreenLoading = true;
                },
                success: function(msg){
                  if(msg.status===4){
                    e.fullscreenLoading = false;
                    e.$message.error(msg.info);
                  }else{

                  }
                    // if(msg.status === 1){
                    //     admin.success(msg.info,'#loginsubmit');
                    //     setTimeout(function(){window.location.href=redirect}, 3000);
                    // }else{
                    //     layer.close(myload);
                    //     admin.error(msg.info,'#loginsubmit');
                    //     formLogin.attr('disabledSubmit','');
                    //     // refreshs();
                    //     $('.yanzheng_img').eq(0).click();
                    //     $('#verify').val('')
                    // }                    
                },
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