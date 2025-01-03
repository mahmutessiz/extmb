<script lang="ts" setup>
import {reactive, ref, computed} from 'vue'
import {RunScript, RunScriptIncrement} from '../../wailsjs/go/main/App'

const data = reactive({
  name: '90',
})

let script= `
	 #!/bin/bash
	 ddcutil setvcp 10 `;

const scriptDecrease=`
	 #!/bin/bash
	 ddcutil setvcp 10 10`;


</script>

<template>
  <main>
    <div id="input" class="input-box">
      <p class="result">Current brightness: {{ data.name }}%, {{ script }}</p>
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" style="width: 200px;" @click="RunScript(scriptDecrease)">10%</button>
      <button class="btn" style="width: 200px;" @click="()=>{
        const script1 = script + data.name
        RunScriptIncrement(script1)
      }">Apply</button>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
