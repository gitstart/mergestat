import { Badge, Button, Input, Label, Panel } from '@mergestat/blocks'
import { CheckIcon, CircleCheckFilledIcon, CircleWarningFilledIcon } from '@mergestat/icons'
import { AuthDetail } from 'src/@types'
import { LINKS_TO, TEST_IDS } from 'src/utils/constants'
import useSetPat from 'src/views/hooks/useSetPat'
import SettedAuth from './setted-auth'

type GitHubAuthProps = {
  idProvider: string
  auth?: AuthDetail
}

const GitHubAuth: React.FC<GitHubAuthProps> = ({ idProvider, auth }: GitHubAuthProps) => {
  const {
    pat,
    showValidation,
    isTokenValid,
    validatePAT,
    changeToken,
    handleSavePAT,
  } = useSetPat(idProvider, auth)

  return (
    <Panel className='rounded-md shadow-sm m-auto'>
      <Panel.Header>
        <div className='w-full flex justify-between'>
          <h4 className='t-h4 mb-0'>Authentication</h4>
          {auth?.credentials &&
            <Badge
              label='Token is set'
              endIcon={<CheckIcon className="t-icon" />}
              variant="success"
            />}
        </div>
      </Panel.Header>

      <Panel.Body className='p-8'>
        {auth && <SettedAuth {...auth} />}

        {!auth && <>
          <span className='mb-6 text-gray-500'>
            In order to access the GitHub API and any private GitHub
            repositories, we need to authenticate with{' '}
            <a
              target='_blank'
              href={LINKS_TO.createPAt}
              className='t-link font-semibold'
              rel='noopener noreferrer'
            >
              a (classic) personal access token
            </a>{' '}
            (PAT). Other authentication methods (such as an OAuth based
            flow) may become available in the future.
          </span>

          <form className='mb-4'>
            <div className='flex flex-col mt-8 mb-5'>
              <Label>GitHub personal access token</Label>
              <div className='flex space-x-3'>
                <Input
                  className='max-w-lg'
                  data-testid={TEST_IDS.patInputText}
                  placeholder='ghp_s0mEsecReTt0k3n'
                  type='password'
                  autoComplete='off'
                  value={pat}
                  onChange={changeToken}
                  onKeyPress={(e) => {
                    if (e.key === 'Enter') {
                      e.preventDefault()
                      handleSavePAT()
                    }
                  }}
                />
                <Button
                  label='Validate'
                  skin='secondary'
                  disabled={pat === ''}
                  data-testid={TEST_IDS.patValidateToken}
                  onClick={validatePAT}
                />

                {showValidation && isTokenValid && (
                  <div className='flex items-center'>
                    <CircleCheckFilledIcon className='t-icon t-icon-success flex-shrink-0' />
                    <p className='text-green-700 ml-1.5'>Token valid</p>
                  </div>
                )}
                {showValidation && !isTokenValid && (
                  <div className='flex items-center'>
                    <CircleWarningFilledIcon className='t-icon t-icon-danger flex-shrink-0' />
                    <p className='text-red-700 ml-1.5'>Token invalid</p>
                  </div>
                )}
              </div>
            </div>
          </form>

          <Button
            label='Save'
            disabled={pat === ''}
            data-testid={TEST_IDS.patSetToken}
            onClick={handleSavePAT}
          />
        </>}
      </Panel.Body>
    </Panel>
  )
}

export default GitHubAuth
