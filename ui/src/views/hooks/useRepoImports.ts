import { useQuery } from '@apollo/client'
import { useEffect, useState } from 'react'
import { RepoImportData } from 'src/@types'
import { GetRepoImportsQuery } from 'src/api-logic/graphql/generated/schema'
import { GET_REPO_IMPORTS } from 'src/api-logic/graphql/queries/get-repo-imports'
import { mapToImportsData } from 'src/api-logic/mappers/imports'
import { useRepositoriesContext, useRepositoriesSetState } from 'src/state/contexts'

const useRepoImports = (poll = false) => {
  const { setShowRemoveImportModal, setImportToRemove } = useRepositoriesSetState()
  const [{ showRemoveImportModal }] = useRepositoriesContext()

  const [imports, setImports] = useState<RepoImportData[]>([])

  const { loading, data, refetch } = useQuery<GetRepoImportsQuery>(GET_REPO_IMPORTS, {
    ...(poll && { pollInterval: 5000 }),
  })

  useEffect(() => {
    setImports(mapToImportsData(data))
  }, [data])

  const prepareToRemove = (id: string, name: string) => {
    setImportToRemove({ id, name })
    setShowRemoveImportModal(true)
  }

  return { loading, imports, showRemoveImportModal, refetch, prepareToRemove }
}

export default useRepoImports
