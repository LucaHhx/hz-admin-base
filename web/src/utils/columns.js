import { getStructNameColumns } from '@/api/system/sysTableColumns'

export const getColumns = (structName) => {
  return getStructNameColumns({ structName }).then((res) => {
    if (res.code === 0) {
      return res.data
    }
    return []
  })
}

// console.log(getColumns('tTest'))
