import React, { useEffect, useState } from 'react'

import { Checkbox } from '@sourcegraph/wildcard'

import { eventLogger } from '../tracking/eventLogger'

import { HAS_DISMISSED_TOAST_STORAGE_KEY, HAS_PERMANENTLY_DISMISSED_TOAST_STORAGE_KEY } from './constants'
import { SurveyRatingRadio } from './SurveyRatingRadio'
import { Toast } from './Toast'
import { getDaysActiveCount } from './util'

const hasBeen30DaysSinceLastNotification = (): boolean => getDaysActiveCount() % 30 === 3

/**
 * Show a toast notification if:
 * 1. User has not recently dismissed the notification
 * 2. User has not permanently dismissed the notification
 * 3. User has been active for 3 days OR has been 30 days since they were last shown the notification
 */
const shouldShowToast = (): boolean =>
    localStorage.getItem(HAS_PERMANENTLY_DISMISSED_TOAST_STORAGE_KEY) !== 'true' &&
    localStorage.getItem(HAS_DISMISSED_TOAST_STORAGE_KEY) !== 'true' &&
    hasBeen30DaysSinceLastNotification()

interface SurveyToastProps {
    /**
     * For Storybook only
     */
    forceVisible?: boolean
}

export const SurveyToast: React.FunctionComponent<SurveyToastProps> = ({ forceVisible }) => {
    const daysActive = getDaysActiveCount()
    const [visible, setVisible] = useState(forceVisible || shouldShowToast())

    useEffect(() => {
        if (visible) {
            eventLogger.log('SurveyReminderViewed')
        }
    }, [visible])

    useEffect(() => {
        if (daysActive % 30 === 0) {
            // Reset toast dismissal 3 days before it will be shown
            localStorage.setItem(HAS_DISMISSED_TOAST_STORAGE_KEY, 'false')
        }
    }, [daysActive])

    const handleDismiss = (): void => {
        localStorage.setItem(HAS_DISMISSED_TOAST_STORAGE_KEY, 'true')
        setVisible(false)
    }

    if (!visible) {
        return null
    }

    return (
        <Toast
            title="Tell us what you think"
            subtitle={
                <span id="survey-toast-scores">How likely is it that you would recommend Sourcegraph to a friend?</span>
            }
            cta={
                <SurveyRatingRadio
                    onChange={handleDismiss}
                    openSurveyInNewTab={true}
                    ariaLabelledby="survey-toast-scores"
                />
            }
            footer={
                <Checkbox
                    id="survey-toast-refuse"
                    label="Don't show this again"
                    onChange={event =>
                        localStorage.setItem(HAS_PERMANENTLY_DISMISSED_TOAST_STORAGE_KEY, String(event.target.checked))
                    }
                />
            }
            onDismiss={handleDismiss}
        />
    )
}
