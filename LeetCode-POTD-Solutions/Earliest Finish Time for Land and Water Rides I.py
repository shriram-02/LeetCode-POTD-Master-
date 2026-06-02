class Solution(object):
    def earliestFinishTime(self, landStartTime, landDuration,
                           waterStartTime, waterDuration):

        ans = float('inf')

        for i in range(len(landStartTime)):
            for j in range(len(waterStartTime)):

                # Land -> Water
                land_end = landStartTime[i] + landDuration[i]
                finish1 = max(land_end, waterStartTime[j]) + waterDuration[j]

                # Water -> Land
                water_end = waterStartTime[j] + waterDuration[j]
                finish2 = max(water_end, landStartTime[i]) + landDuration[i]

                ans = min(ans, finish1, finish2)

        return ans