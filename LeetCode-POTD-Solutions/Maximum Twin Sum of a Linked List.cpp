class Solution {
public:
    int pairSum(ListNode* head) {
        ListNode *slow = head, *fast = head;

        while (fast && fast->next) {
            slow = slow->next;
            fast = fast->next->next;
        }

        ListNode *prev = nullptr, *curr = slow;
        while (curr) {
            ListNode *nextNode = curr->next;
            curr->next = prev;
            prev = curr;
            curr = nextNode;
        }

        int ans = 0;
        ListNode *p1 = head, *p2 = prev;

        while (p2) {
            ans = max(ans, p1->val + p2->val);
            p1 = p1->next;
            p2 = p2->next;
        }

        return ans;
    }
};