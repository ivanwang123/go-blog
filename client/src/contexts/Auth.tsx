import axios from "axios";
import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";
import { UserType } from "src/types/UserType";

type ContextType = {
  isAuthenticated: boolean;
  me: UserType | null;
  getMe: () => void;
};

const AuthContext = createContext<ContextType>({
  isAuthenticated: false,
  me: null,
  getMe: () => {},
});

export const useAuthContext = () => useContext(AuthContext);

function AuthProvider({ children }: { children: ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [me, setMe] = useState<UserType | null>(null);
  const [loading, setLoading] = useState<boolean>(true);

  const getMe = async () => {
    console.log("GET ME");
    try {
      const res = await axios.get("/user/me");
      setIsAuthenticated(true);
      setMe(res.data.user);
    } catch (e) {
      setIsAuthenticated(false);
      setMe(null);
    }
    setLoading(false);
  };

  useEffect(() => {
    getMe();
  }, []);

  const value = useMemo(
    () => ({
      isAuthenticated,
      me,
      getMe,
    }),
    [isAuthenticated, me]
  );

  return (
    <>
      {loading ? (
        <h1>Loading</h1>
      ) : (
        <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
      )}
    </>
  );
}

export default AuthProvider;
