import {
  createProSearchForm,
  useNDataTable,
  type UseNDataTableData,
  type UseNDataTableParams,
  type UseNDataTableService,
} from 'pro-naive-ui'

type InitValues = Record<string,any>

export function useCrud<T, R>(
  service: UseNDataTableService<UseNDataTableData, UseNDataTableParams>,
  initValues?:InitValues
) {
  const loading = ref(false)
  const searchForm = createProSearchForm({
    defaultCollapsed: true,
    onReset: () => {},
    onSubmit: () => {
      loading.value = true
      setTimeout(() => {
        loading.value = false
      }, 1500)
    },
  })
  const { table, search } = useNDataTable(service, {
    form: searchForm,
  })

  return {
    loading,
    searchForm,
    search,
    table,
  }
}
