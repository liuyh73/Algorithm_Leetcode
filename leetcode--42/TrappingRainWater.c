int trap(int* height, int heightSize) {
    if(heightSize == 0) return 0;
    int *head=height, *tail=height+(heightSize-1);
    // minnEdge denote the left or right edge.
    int unitsOfRain=0, minnEdge=-1, minnValue;
    while(head != tail){
        if(*head < *tail){
            // only when the minnEdge changes, the minnValue will change
            if(minnEdge == 1 || minnEdge == -1){
                minnValue=*head;
                minnEdge = 0;
            }
            head++;
            if(*head < minnValue) 
                unitsOfRain += minnValue - *head;
            else
                minnValue = *head;
        }
        else {
            if(minnEdge == 0 || minnEdge == -1){
                minnValue = *tail;
                minnEdge = 1;
            }
            tail--;
            if(*tail < minnValue)
                unitsOfRain += minnValue - *tail;
            else
                minnValue = *tail;
        }
    }
    return unitsOfRain;
}