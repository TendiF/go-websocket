import { createSlice } from '@reduxjs/toolkit';
import axios from './utils/axios'
export const userReducer = createSlice({
  name: 'counter',
  initialState: {
    user: {},
  },
  reducers: {
    storeUser: (state, action) => {
      state.user = action.payload;
    },
  },
});

export const { storeUser } = userReducer.actions;

export const doLogin = data => dispatch => {
  axios.post('/user/login', data)
  .then(res => {
    console.log(res)
  })
  .catch(err => {

  })
  // dispatch(storeUser(data));
};
export const getUser = state => state.counter.value;

export default userReducer.reducer;
