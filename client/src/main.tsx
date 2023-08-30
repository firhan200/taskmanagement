import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import 'react-date-picker/dist/DatePicker.css';
import 'react-calendar/dist/Calendar.css';
import 'react-time-picker/dist/TimePicker.css';
import 'react-clock/dist/Clock.css';

import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'
import AppRouter from './AppRouter'

// Create a client
const queryClient = new QueryClient()

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.Fragment>
    <QueryClientProvider client={queryClient}>
      <AppRouter />
    </QueryClientProvider>
  </React.Fragment>,
)
