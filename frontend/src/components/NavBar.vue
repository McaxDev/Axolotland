<script setup lang="ts">
import {ref} from 'vue';
import type { BaseColorVariant } from 'bootstrap-vue-next';
import {useRouter} from 'vue-router';
import {wikiMeta} from '@/utils/wiki';
import {useI18n} from 'vue-i18n';
import {langDisplayName} from '@/i18n/metadata';

const router = useRouter()
const navVariant = ref<keyof BaseColorVariant>('light')
const navColorMode = ref('light')
const i18n = useI18n()

// 顶栏外部链接列表
const navDropdownLink: Record<string, string> = {
	bbs: 'bbs.mcax.cn',
	map: 'map.mcax.cn',
	log: 'log.mcax.cn',
	cloud: 'cloud.mcax.cn',
}

// 修改主题相关
const themes: Record<keyof BaseColorVariant, boolean> = {
	primary: true,
	secondary: true,
	success: true,
	info: false,
	warning: true,
	danger: false,
	light: false,
	dark: true,
}

function setTheme(theme: keyof BaseColorVariant) {
  document.documentElement.setAttribute('theme', theme)
	navVariant.value = theme
	navColorMode.value = themes[theme] === true ? 'dark' : 'light'
}
</script>

<template>
  <b-navbar toggleable="md" :variant="navVariant" v-b-color-mode="navColorMode" fixed="top" class="navbar">

		<b-navbar-brand to="/">
      <img src="./assets/axo-logo.png" height="30px" alt="" />
    </b-navbar-brand>

    <b-navbar-toggle target="nav-collapse" />

    <b-collapse is-nav id="nav-collapse">

			<!-- 顶部页面导航 -->
      <b-navbar-nav>

				<!-- 主页 -->
				<b-nav-item to="/">
					<img src="./assets/mcicon/mainpage.png" height="20px" alt="" />
          {{ $t('navbar.routes.mainpage') }}
        </b-nav-item>

				<!-- 维基百科 -->
				<b-nav-item-dropdown right>
					<template #button-content>
						<span @click="router.push('/wiki/index.html')">
						<img src="./assets/mcicon/wiki.png" height="20px" alt="" />
						{{ $t('navbar.routes.wiki') }}
						</span>
					</template>
					<b-dropdown-item v-for="(value, key) in wikiMeta" :key="key"
						:to="`/wiki/${key}.html`"
					>{{ value }}</b-dropdown-item>
				</b-nav-item-dropdown>

				<!-- 画廊 -->
				<b-nav-item to="/gallery">
					<img src="./assets/mcicon/gallery.png" height="20px" alt="" />
          {{ $t('navbar.routes.gallery') }}
        </b-nav-item>

				<!-- 仪表板 -->
				<b-nav-item to="/dash">
					<img src="./assets/mcicon/dash.png" height="20px" alt="" />
          {{ $t('navbar.routes.dash') }}
        </b-nav-item>
      </b-navbar-nav>

      <b-navbar-nav class="ms-auto">

				<!-- 相关链接 -->
				<b-nav-item-dropdown right>
					<template #button-content>
						{{ $t('navbar.external-link') }}
					</template>
					<b-dropdown-item v-for="(value, key) in navDropdownLink"
						:key="key" :href="'https://'+value"
					>
					  {{ $t('navbar.external-links.'+key) }}
					</b-dropdown-item>	
				</b-nav-item-dropdown>

        <!-- 多语言切换按钮 -->
        <b-nav-item-dropdown no-caret right>
          <template #button-content>
            <i class="bi bi-translate" />
          </template>
          <b-dropdown-item
            v-for="(value, key) in langDisplayName"
            :key="key"
            @click="i18n.locale.value = key"
						>{{ value }}
					</b-dropdown-item>
        </b-nav-item-dropdown>

        <!-- 多主题切换按钮 -->
        <b-nav-item-dropdown no-caret right>
          <template #button-content>
            <i class="bi bi-palette" />
          </template>
          <b-dropdown-item v-for="(_, key) in themes" :key="key" @click="setTheme(key)">
						<i class="bi bi-palette-fill" :style="`color: var(--bs-${key})`" />
					</b-dropdown-item>
        </b-nav-item-dropdown>

				<!-- Github -->
				<b-nav-item href="https://github.com/McaxDev">
					<i class="bi bi-github" />
				</b-nav-item>

				<!-- 个人中心 -->
				<b-nav-item to="/user">
					<i class="bi bi-person" />
				</b-nav-item>
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>

</template>
