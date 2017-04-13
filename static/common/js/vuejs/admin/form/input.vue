<style>
body{font-family: Helvetica Neue,Helvetica,PingFang SC,Hiragino Sans GB,Microsoft YaHei,SimSun,sans-serif;font-weight: 400;}
input,textarea{font-family: inherit;font-size: inherit;}
input::input-placeholder { line-height: normal; }
input::-ms-input-placeholder{line-height: normal;}
input::-webkit-input-placeholder{line-height: normal;}
.hm_input{
	display: inline-block;
    vertical-align: top;
    margin: 10px 5px;
    position: relative;
    width: 180px;
    font-size: 14px;
    height: 36px;
}
.error{position: absolute;top:44px;z-index: 100;display:block;
	background: #324057;padding: 5px 5px;border-radius: 4px;
	
}
.error:before{
	position: absolute;
    left: 13px;
    top: -8px;
    width: 0;
    height: 0;
    content: "";
    border-width: 0px 8px 8px 8px;
    border-style: solid;
    border-color: #324057 transparent;
    -webkit-transition: transform .25s;
       -moz-transition: transform .25s;
        -ms-transition: transform .25s;
         -o-transition: transform .25s;
            transition: transform .25s;
}
.error span{color: white;}
/*.errors:hover::before { content: "\5B"; left: -20px; }*/
.hm_textarea{width: 180px;}
.hm_input_group{min-width: 260px;}
.hm_input_min{font-size: 14px;}
.hm_input_small{font-size: 16px;}
.hm_input_big{font-size: 18px;}
.hm_input_max{font-size: 24px;}
.hm_input_inner{display: block;padding: 3px 10px;box-sizing: border-box;width: 100%;
    height: 36px;color: #1f2d3d;background-color: #fff;background-image: none;
    border: 1px solid #c0ccda;border-radius: 4px;transition: border-color .2s cubic-bezier(.645,.045,.355,1);-webkit-transition: border-color .2s cubic-bezier(.645,.045,.355,1);outline: none;line-height: normal;}
.hm_input_inner:focus{outline: none; border-color: #20a0ff;}
.hm_input_min .hm_input_inner{height:36px;font-size: 14px;}
.hm_input_small .hm_input_inner{height:40px;font-size: 20px;}
.hm_input_big .hm_input_inner{height: 44px;font-size: 24px;}
.hm_input_max .hm_input_inner{height: 48px;font-size:28px;}
.hm_input_min .error{top:44px;}
.hm_input_small .error{top:48px;}
.hm_input_big .error{top:52px;}
.hm_input_max .error{top:56px;}
.hm_input i.fa{
	position: absolute;	
	top:5%;
    width: 36px;
    height: 100%;
    text-align: center;
    color: #c0ccda;}

.hm_input i.fa:after {
    content: '';
    height: 100%;
    width: 0;
    display: inline-block;
    vertical-align: middle;
}
.hm_input i.clear{position: relative;top:0.1em;width: 100%;}
.hm_input i.left+.hm_input_inner{padding-left: 36px;}
.hm_input i.right+.hm_input_inner{padding-right: 36px;}
.hm_input i.left{left: 0;}
.hm_input i.right{right: 0;}

.hm_input_group_left, .hm_input_group_right {
    background-color: #f9fafc;
    color: #99a9bf;
    vertical-align: middle;
    display: table-cell;
    position: relative;
    border: 1px solid #c0ccda;
    border-radius: 4px;
    padding: 0 10px;
    width: 1%;
    white-space: nowrap;
}
.hm_input_group_left{border-right: 0;border-top-right-radius: 0; border-bottom-right-radius: 0;}
.hm_input_group_right {border-left: 0;border-top-left-radius: 0;
    border-bottom-left-radius: 0;
}
.hm_input_group {
    display: table;
    border-collapse: separate;
}
.hm_input_group>.hm_input_inner {
    vertical-align: middle;
    display: table-cell;
}
.hm_input_group .hm_input_inner:last-child {
    border-top-left-radius: 0;
    border-bottom-left-radius: 0;
}
.hm_input_group .hm_input_inner:first-child {
    border-top-right-radius: 0;
    border-bottom-right-radius: 0;
}
.hm_input_group .hm_input_inner:not(:first-child):not(:last-child) {
    border-radius: 0;
}
</style>
<template>
	<div :style="{ width: width + 'px' }" :class="[
		type === 'textarea' ? 'hm_textarea' : 'hm_input',
		size ? 'hm_input_' + size : '',	
		{
	      'is_disabled': disabled,
	      'hm_input_group': $slots.left || $slots.right
	    }
	]">
		<template v-if="type !== 'textarea'">
			<!-- 前置元素 -->
		    <div :style="{ width: slotleft + 'px' }" class="hm_input_group_left" v-if="$slots.left">
		        <slot name="left"></slot>
		    </div>
			<!-- input 图标 -->
      		<i class="fa" :class="[icon ? 'fa-' + icon : '']" v-if="icon"></i>
      		<!-- <i class="el-input__icon el-icon-loading" v-if="validating"></i> -->
			<input 
				v-if="type === 'text'"
				class="hm_input_inner"
				v-model="inputValue"
				type="text"			
				:class="[validata === '' ? '' : validata]"	
				:name="name"
				:placeholder="placeholder"
				:disabled="disabled"
				:maxlength="maxlength"
        		:minlength="minlength"
        		:readonly="readonly"
        		:autocomplete="autoComplete"
        		ref="input"
        		@focus="handleFocus"
        		@blur="handleBlur"
			>
			<input 
				v-if="type === 'password'"
				class="hm_input_inner"
				v-model="inputValue"
				type="password"				
				:name="name"
				:placeholder="placeholder"
				:disabled="disabled"
				:maxlength="maxlength"
        		:minlength="minlength"
        		:readonly="readonly"
        		:autocomplete="autoComplete"
        		ref="input"
        		@focus="handleFocus"
        		@blur="handleBlur"
			>
			<!-- <div class="errors"></div> -->
			<div class="hm_input_group_right" v-if="$slots.right">
		        <slot name="right"></slot>
		    </div>
			<!-- input 右侧图标 -->
      		<!-- <i class="el-input__icon" :class="[icon ? 'el-icon-' + icon : '']" v-if="icon"></i>
      		<i class="el-input__icon el-icon-loading" v-if="validating"></i> -->
		</template>
		<textarea 
			v-else
			class="hm_input_inner"
			v-model="inputValue"			
			ref="textarea"			
			:name="name"
			:placeholder="placeholder"
			:disabled="disabled"
			:maxlength="maxlength"
    		:minlength="minlength"
    		:readonly="readonly"
    		:style="textareaStyle"
    		:rows="rows"
    		:autocomplete="autoComplete"
    		@focus="handleFocus"
    		@blur="handleBlur">			
		</textarea>
	</div>
</template>
<script>
import calcTextareaHeight from './calcTextareaHeight';
import validate from './validate.js';
export default{
	// name: 'HmInput',
	props:{
		value: [String, Number],
		type: {//类型text,password,textarea之类的
	        type: String,
	        default: 'text'
	    },
	    validata: {//输入验证的数据
	        type: String,
	        default: 'text'
	    },
	    placeholder: {//input提示信息,默认为空
	        type: String,
	        default: ''
	    },
	    icon: {//图标
	        type: String,
	        default: ''
	    },
	    size: {//长度大小
        	type: String,
        	default: 'min'
      	},
      	width: {//长度大小
        	type: String,
        	default: '200'
      	},
      	slotleft: {//长度大小
        	type: String,
        	default: '100'
      	},
      	name: {//form提交用到的name
	        type: String,
	        default: ''
	    },
	    number: {
        	type: Boolean,
	        default: false
	    },
	    disabled: {//默认为不禁用
	        type: Boolean,
	        default: false
	    },
	    autoComplete: {//输入的字段自动完成,默认为开启:on,关闭off
	        type: String,
	        default: 'on'
	    },
	    autosize: {
	        type: [Boolean, Object],
	        default: false
      	},
	    readonly: {//默认为false:关闭只读,true只读
	        type: Boolean,
	        default: false
	    },
	    rows: {
	        type: Number,
	        default: 2
	    },
	    maxlength: Number,
      	minlength: Number

	},
	methods: {
		handleFocus(e){
			// this.$emit('onblur', this.inputValue);
			// alert('focus');
		},
		handleBlur(e){
			// alert('blur');
		},
		inputSelect() {
        	this.$refs.input.select();
      	},
      	resizeTextarea() {
	        var { autosize, type } = this;
	        if (!autosize || type !== 'textarea') {
	          return;
	        }
	        const minRows = autosize ? autosize.minRows : null;
	        const maxRows = autosize ? autosize.maxRows : null;
	        this.textareaStyle = calcTextareaHeight(this.$refs.textarea, minRows, maxRows);
	    },

	},
	created() {
      this.$on('inputSelect', this.inputSelect);
    },
	mounted() {
		// console.info('1111');
	},
	data(){
		return{
			inputValue:this.value,
			textareaStyle: {}
		}
	},
	watch:{		
		'inputValue'(val, oldValue) {
			// this.inputValue = val;
   //     	    this.resizeTextarea();
   //     	    this.$http.post('/someUrl').then((response) => {
			//     // success callback
			//   }, (response) => {
			//     // error callback
			// });
   //     	    validate(this.$refs.input);
       	    // $(this.$refs.input).click(function(){
       	    // 	alert('111');
       	    // });

		}
	}
}
</script>