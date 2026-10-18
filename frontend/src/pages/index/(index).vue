<template>
    <q-page class="q-pa-md bg-grey-1">
        <!-- Header com título -->
        <div class="row items-center justify-between q-mb-md">
            <div>
                <h1 class="text-h5 text-bold q-my-none text-primary">Dashboard Geral</h1>
                <div class="text-caption text-grey-7">Visão geral e métricas do sistema</div>
            </div>
            <q-btn icon="refresh" label="Atualizar" color="primary" outline dense class="q-px-sm" />
        </div>

        <!-- Cards de Métricas (KPIs) -->
        <div class="row q-col-gutter-md q-mb-md">
            <div class="col-12 col-sm-6 col-md-3">
                <q-card class="shadow-1">
                    <q-card-section class="row items-center justify-between">
                        <div>
                            <div class="text-grey-7 text-caption">Total de Vendas</div>
                            <div class="text-h6 text-bold">R$ 48.290,00</div>
                        </div>
                        <q-avatar icon="payments" color="green-1" text-color="green" />
                    </q-card-section>
                </q-card>
            </div>

            <div class="col-12 col-sm-6 col-md-3">
                <q-card class="shadow-1">
                    <q-card-section class="row items-center justify-between">
                        <div>
                            <div class="text-grey-7 text-caption">Novos Usuários</div>
                            <div class="text-h6 text-bold">+1.240</div>
                        </div>
                        <q-avatar icon="people" color="blue-1" text-color="blue" />
                    </q-card-section>
                </q-card>
            </div>

            <div class="col-12 col-sm-6 col-md-3">
                <q-card class="shadow-1">
                    <q-card-section class="row items-center justify-between">
                        <div>
                            <div class="text-grey-7 text-caption">Acessos Hoje</div>
                            <div class="text-h6 text-bold">8.420</div>
                        </div>
                        <q-avatar icon="trending_up" color="purple-1" text-color="purple" />
                    </q-card-section>
                </q-card>
            </div>

            <div class="col-12 col-sm-6 col-md-3">
                <q-card class="shadow-1">
                    <q-card-section class="row items-center justify-between">
                        <div>
                            <div class="text-grey-7 text-caption">Taxa de Conversão</div>
                            <div class="text-h6 text-bold">3,85%</div>
                        </div>
                        <q-avatar icon="show_chart" color="orange-1" text-color="orange" />
                    </q-card-section>
                </q-card>
            </div>
        </div>

        <!-- Seção dos Gráficos -->
        <div class="row q-col-gutter-md">
            <!-- Gráfico de Linhas: Desempenho de Vendas -->
            <div class="col-12 col-md-8">
                <q-card class="shadow-1">
                    <q-card-section>
                        <div class="text-subtitle1 text-bold">Desempenho de Vendas (2026)</div>
                    </q-card-section>
                    <q-card-section>
                        <v-chart class="chart" :option="salesChartOption" autoresize />
                    </q-card-section>
                </q-card>
            </div>

            <!-- Gráfico de Rosca: Distribuição de Categorias -->
            <div class="col-12 col-md-4">
                <q-card class="shadow-1">
                    <q-card-section>
                        <div class="text-subtitle1 text-bold">Vendas por Categoria</div>
                    </q-card-section>
                    <q-card-section>
                        <v-chart class="chart" :option="categoryChartOption" autoresize />
                    </q-card-section>
                </q-card>
            </div>
        </div>
    </q-page>
</template>

<script setup>
import { ref } from 'vue'
import VChart from 'vue-echarts'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart, PieChart } from 'echarts/charts'
import {
    GridComponent,
    TooltipComponent,
    LegendComponent,
    TitleComponent
} from 'echarts/components'

// Registra os módulos necessários do ECharts para manter o bundle leve
use([
    CanvasRenderer,
    LineChart,
    PieChart,
    GridComponent,
    TooltipComponent,
    LegendComponent,
    TitleComponent
])

// Configuração do Gráfico de Linhas
const salesChartOption = ref({
    tooltip: {
        trigger: 'axis'
    },
    grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true
    },
    xAxis: {
        type: 'category',
        boundaryGap: false,
        data: ['Jan', 'Fev', 'Mar', 'Abr', 'Mai', 'Jun', 'Jul', 'Ago', 'Set', 'Out']
    },
    yAxis: {
        type: 'value',
        axisLabel: {
            formatter: 'R$ {value}'
        }
    },
    series: [
        {
            name: 'Vendas (R$)',
            type: 'line',
            smooth: true,
            data: [12000, 19000, 15000, 22000, 28000, 26000, 31000, 38000, 42000, 48290],
            areaStyle: {
                opacity: 0.15
            },
            itemStyle: {
                color: '#1976D2'
            }
        }
    ]
})

// Configuração do Gráfico de Rosca (Pie Chart)
const categoryChartOption = ref({
    tooltip: {
        trigger: 'item'
    },
    legend: {
        bottom: '0',
        left: 'center'
    },
    series: [
        {
            name: 'Categorias',
            type: 'pie',
            radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            itemStyle: {
                borderRadius: 8,
                borderColor: '#fff',
                borderWidth: 2
            },
            label: {
                show: false
            },
            data: [
                { value: 1048, name: 'Assinaturas' },
                { value: 735, name: 'Licenças' },
                { value: 580, name: 'Serviços' },
                { value: 484, name: 'Outros' }
            ]
        }
    ]
})
</script>

<style scoped>
.chart {
    height: 320px;
    width: 100%;
}
</style>