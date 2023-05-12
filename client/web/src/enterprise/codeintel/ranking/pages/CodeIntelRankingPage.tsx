import { FunctionComponent, useEffect } from 'react'

import { Timestamp } from '@sourcegraph/branded/src/components/Timestamp'
import { TelemetryProps, TelemetryService } from '@sourcegraph/shared/src/telemetry/telemetryService'
import { ErrorAlert, LoadingSpinner, PageHeader } from '@sourcegraph/wildcard'

import { useRankingSummary as defaultUseRankingSummary } from './backend'

export interface CodeIntelRankingPageProps extends TelemetryProps {
    useRankingSummary?: typeof defaultUseRankingSummary
    telemetryService: TelemetryService
}

export const CodeIntelRankingPage: FunctionComponent<CodeIntelRankingPageProps> = ({
    useRankingSummary = defaultUseRankingSummary,
    telemetryService,
}) => {
    useEffect(() => telemetryService.logViewEvent('CodeIntelRankingPage'), [telemetryService])

    const { data, loading, error } = useRankingSummary({})

    if (loading && !data) {
        return <LoadingSpinner />
    }

    if (error) {
        return <ErrorAlert prefix="Failed to load code intelligence summary for repository" error={error} />
    }

    return (
        <>
            <PageHeader
                headingElement="h2"
                path={[
                    {
                        text: <>Ranking calculation history</>,
                    },
                ]}
                description="View the history of ranking calculation."
                className="mb-3"
            />
            {data && data.rankingSummary.map(summary => <Summary key={summary.graphKey} summary={summary} />)}
        </>
    )
}

interface Summary {
    graphKey: string
    pathMapperProgress: Progress
    referenceMapperProgress: Progress
    reducerProgress: Progress | null
}

interface Progress {
    startedAt: string
    completedAt: string | null
    processed: number
    total: number
}

interface SummaryProps {
    summary: Summary
}

const Summary: FunctionComponent<SummaryProps> = ({ summary }) => (
    <div>
        <strong>Graph key: {summary.graphKey}</strong>

        <ul>
            <li>
                Path mapper: <Progress progress={summary.pathMapperProgress} />
            </li>

            <li>
                Reference mapper: <Progress progress={summary.referenceMapperProgress} />
            </li>

            {summary.reducerProgress && (
                <li>
                    Reducer: <Progress progress={summary.reducerProgress} />
                </li>
            )}
        </ul>
    </div>
)

interface ProgressProps {
    progress: Progress
}

const Progress: FunctionComponent<ProgressProps> = ({ progress }) => (
    <span>
        {progress.processed} of {progress.total} records processed; started <Timestamp date={progress.startedAt} />{' '}
        {progress.completedAt && (
            <>
                and completed <Timestamp date={progress.completedAt} />
            </>
        )}
    </span>
)
