<script setup lang="ts">
import { apiClient, type Resp } from '@/utils/request'
import { onMounted, ref } from 'vue'
import * as echarts from 'echarts'
import { fetchNginxFile } from '@/utils/nginx'

interface Chart {
  name: string
  display: string
}

const charts = ref<Chart[]>([])

onMounted(async () => {
	charts.value = await fetchNginxFile<Chart[]>('/dash/statslist.json')
	for (const chart of charts.value) {
		const data = await apiClient.get<Resp<{
  		Member: string
  		Score: number
		}[]>>(`/dash/${chart.name}`)
		const sorted = data.data.data.sort((a, b) => a.Score - b.Score)
		const instance = echarts.init(document.getElementById(chart.name))
		instance.setOption({
			title: {
				text: chart.display
			},
			tooltip: {
				trigger: 'axis'
			},
			xAxis: {
				type: 'value'
			},
			yAxis: {
				type: 'category',
				data: sorted.map(item => item.Member),
				axisLabel: {
					formatter: function(value: string) {
						return value.length > 8 ? value.substring(0, 8) : value;
					},
					rotate: 45,
				}
			},
			series: [
				{
					name: chart.display,
					type: 'bar',
					data: sorted.map(item => item.Score)
				}
			]
		})
		window.addEventListener('resize', () => {
			instance.resize()
		})
	}
})

</script>

<template>
  <div class="margin-nav py-3">
    <b-container class="p-4 rounded shadow">
      <b-row>
        <b-col v-for="item in charts" :key="item.name" 
					cols="12" md="6"
				>
          <div :id="item.name" class="chart-container"></div>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
}
</style>
