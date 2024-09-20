<script setup lang="ts">
import {ref} from 'vue';

const mode = ref<'login' | 'signup' | 'forgot'>('login')
const email = ref<string>('')
const captchaId = ref<string>('')
const captchaValue = ref<string>('')
</script>

<template>
	<b-container class="d-flex justify-content-center align-items-center vh-100">
		<b-row class="w-100">
			<b-col class="align-items-center justify-content-center d-none d-md-flex" md="6">
			</b-col>
			<b-col cols="12" md="6" class="d-flex justify-content-center">

			</b-col>
		</b-row>
	</b-container>
</template>
