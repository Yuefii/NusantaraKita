import { createContext, useContext, useReducer, type ReactNode } from 'react';

// State type
export interface MapState {
  isShowEndpoints: boolean;
  selected: {
    province: ProvinsiApi | null;
    kabKota: KabKotaApi | null;
    kecamatan: KecamatanApi | null;
    desaKel: DesaKelApi | null;
  };
}

// Action type
type MapAction =
  | { type: 'TOGGLE_ENDPOINTS'; payload: boolean }
  | { type: 'SET_PROVINCE'; payload: ProvinsiApi }
  | { type: 'SET_KABKOTA'; payload: KabKotaApi | null }
  | { type: 'SET_KECAMATAN'; payload: KecamatanApi | null }
  | { type: 'SET_DESA'; payload: DesaKelApi | null }
  | { type: 'RESET_MAP' };

// Initial state
const initialState: MapState = {
  isShowEndpoints: false,
  selected: {
    province: null,
    kabKota: null,
    kecamatan: null,
    desaKel: null,
  },
};

// Reducer function
function mapReducer(state: MapState, action: MapAction): MapState {
  switch (action.type) {
    case 'TOGGLE_ENDPOINTS':
      return { ...state, isShowEndpoints: action.payload };

    case 'SET_PROVINCE':
      return {
        ...state,
        selected: { ...state.selected, province: action.payload },
      };
    case 'SET_KABKOTA':
      return {
        ...state,
        selected: { ...state.selected, kabKota: action.payload },
      };
    case 'SET_KECAMATAN':
      return {
        ...state,
        selected: { ...state.selected, kecamatan: action.payload },
      };
    case 'SET_DESA':
      return {
        ...state,
        selected: { ...state.selected, desaKel: action.payload },
      };

    case 'RESET_MAP':
      return initialState;

    default:
      return state;
  }
}

// Context type
interface OverviewContextType {
  state: MapState;
  dispatch: React.Dispatch<MapAction>;
}

// Create context
const OverviewContext = createContext<OverviewContextType | undefined>(
  undefined,
);

// Provider component
export const OverviewProvider = ({ children }: { children: ReactNode }) => {
  const [state, dispatch] = useReducer(mapReducer, initialState);

  return (
    <OverviewContext.Provider value={{ state, dispatch }}>
      {children}
    </OverviewContext.Provider>
  );
};

// Custom hook
export const useOverview = (): OverviewContextType => {
  const context = useContext(OverviewContext);
  if (!context)
    throw new Error('useOverview must be used within an OverviewProvider');

  return context;
};
