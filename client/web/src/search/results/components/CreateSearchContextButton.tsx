import React from 'react'

import { ButtonLink } from '@sourcegraph/shared/src/components/LinkOrButton'
import { omitFilter } from '@sourcegraph/shared/src/search/query/transformer'
import { FilterKind, findFilter } from '@sourcegraph/shared/src/search/query/validate'
import { FilterType } from '@sourcegraph/shared/src/search/query/filters'

import { AuthenticatedUser } from '../../../auth'
import MagnifyIcon from 'mdi-react/MagnifyIcon'

interface CreateSearchContextButtonProps {
    /** Search query string. */
    query?: string

    /** The currently authenticated user or null. */
    authenticatedUser: Pick<AuthenticatedUser, 'id'> | null
}

export const CreateSearchContextButton: React.FunctionComponent<CreateSearchContextButtonProps> = props => {
    if (!window.context.experimentalFeatures['search.contexts.query'] || !props.query || !props.authenticatedUser) {
        return null
    }

    const contextFilter = findFilter(props.query, FilterType.context, FilterKind.Global)
    if (!contextFilter || contextFilter.value?.value != 'global') {
        return null
    }

    const query = omitFilter(props.query, contextFilter)
    const searchParameters = new URLSearchParams()
    searchParameters.set('q', query)
    const toURL = `/contexts/new?${searchParameters.toString()}`

    return (
        <li data-tooltip="Create search context based on this query" data-delay={10000} className="nav-item mr-2">
            <ButtonLink to={toURL} className="btn btn-sm btn-outline-secondary text-decoration-none">
                <MagnifyIcon className="icon-inline mr-1" />
                Create context
            </ButtonLink>
        </li>
    )
}
