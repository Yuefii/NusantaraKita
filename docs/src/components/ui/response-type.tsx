'use client';

import React, {
  createContext,
  useContext,
  useReducer,
  type Dispatch,
  type ReactNode,
} from 'react';

import { cn } from '@/lib/utils';
import { CiViewTable } from 'react-icons/ci';
import { VscJson } from 'react-icons/vsc';

// Type value yang dipakai context
export type ResponseTypeValue = 'TABLE' | 'JSON';

// State type
interface State {
  dataType: ResponseTypeValue;
}

// Action type
type Action = { type: 'SET_RESPONSE_TYPE'; payload: ResponseTypeValue };

// Initial state
const initialState: State = {
  dataType: 'TABLE',
};

// Reducer function
function responseTypeReducer(state: State, action: Action): State {
  switch (action.type) {
    case 'SET_RESPONSE_TYPE':
      return { ...state, dataType: action.payload };
    default:
      return state;
  }
}

// Context types
interface ResponseTypeContextProps {
  state: State;
  dispatch: Dispatch<Action>;
}

// Create context
const ResponseTypeContext = createContext<ResponseTypeContextProps | undefined>(
  undefined,
);

// Provider component
export const ResponseTypeProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [state, dispatch] = useReducer(responseTypeReducer, initialState);

  return (
    <ResponseTypeContext.Provider value={{ state, dispatch }}>
      {children}
    </ResponseTypeContext.Provider>
  );
};

// Custom hook buat akses context
export function useResponseType() {
  const context = useContext(ResponseTypeContext);
  if (!context) {
    throw new Error(
      'useResponseType must be used within a ResponseTypeProvider',
    );
  }
  return context;
}

export default function ResponseTypeSwitch() {
  const { state, dispatch } = useResponseType();

  const handleTableClick = () =>
    dispatch({ type: 'SET_RESPONSE_TYPE', payload: 'TABLE' });
  const handleJsonClick = () =>
    dispatch({ type: 'SET_RESPONSE_TYPE', payload: 'JSON' });

  return (
    <div className="border-border bg-muted/60 relative flex w-fit items-center overflow-hidden rounded-lg border p-1">
      {/* Background slider */}
      <div
        className={cn(
          'bg-background absolute left-1 h-8 w-8 rounded-md shadow-xs transition-transform duration-200 ease-in-out',
          state.dataType === 'TABLE' ? 'translate-x-0' : 'translate-x-full',
        )}
      />

      {/* Table Button */}
      <button
        onClick={handleTableClick}
        className={cn(
          'relative z-10 flex h-8 w-8 cursor-pointer items-center justify-center text-lg transition-colors duration-150',
          state.dataType === 'TABLE'
            ? 'text-foreground font-semibold'
            : 'text-muted-foreground hover:text-foreground',
        )}
        aria-label="Table View"
        title="Tampilan Tabel"
      >
        <CiViewTable className="h-5 w-5" />
      </button>

      {/* JSON Button */}
      <button
        onClick={handleJsonClick}
        className={cn(
          'relative z-10 flex h-8 w-8 cursor-pointer items-center justify-center text-lg transition-colors duration-150',
          state.dataType === 'JSON'
            ? 'text-foreground font-semibold'
            : 'text-muted-foreground hover:text-foreground',
        )}
        aria-label="JSON View"
        title="Tampilan JSON"
      >
        <VscJson className="h-5 w-5" />
      </button>
    </div>
  );
}
