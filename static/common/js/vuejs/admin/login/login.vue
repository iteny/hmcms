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
      <el-input type="password" v-model="hmForm.password"></el-input>
    </el-form-item>
    <el-form-item label="验证码" prop="captcha"  style="position: relative;">
        <p><img :src="imgurl" style="position: absolute;top:2px;right: 2px;z-index:999;" /></p>
        <el-input v-model="hmForm.captcha" type="text" :maxlength="15"></el-input>
        <input v-model="hmForm.captcha_id" type="hidden" />
    </el-form-item>
    <el-form-item>
      <el-button id="loginsubmit" type="primary" @click="submitForm('hmForm')" style="width:100%;" v-loading.fullscreen.lock="fullscreenLoading">立即创建</el-button>
      <!-- <el-button @click="resetForm('hmForm')">重置</el-button> -->
    </el-form-item>
  </el-form>
</template>
<script>
  export default {
    props:{
        captchaid:{
            type: String,
            default: 'text'
        },
        imgurl:{
            type: String,
            default: 'text'
        }
    },
    data() {
      return {
        fullscreenLoading: false,
        hmForm: {
          username: '',
          password: '',
          captcha:'',
          captcha_id:'',
        },
        rules: {
          username: [
            { required: true, message: '请输入用户名', trigger: 'blur' },
            { min: 5, max: 15, message: '长度在 5 到 15 个字符', trigger: 'blur' }
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
                data : {username:this.hmForm.username,password:this.hmForm.password,captcha:this.hmForm.captcha,captcha_id:this.captchaid},
                dataType:"json",
                beforeSend: function(){
                    e.fullscreenLoading = true;
                },
                success: function(msg){
                  if(msg.status===4){
                    e.fullscreenLoading = false;
                    e.$message.error(msg.info);
                  }else if(msg.status===1){
                    e.fullscreenLoading = false;
                    e.$message.success(msg.info);
                    setTimeout(function(){window.location.href="/admin/index"}, 3000);
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
