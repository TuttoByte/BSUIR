using System.Collections.Generic;
using AvaloniaApplication7.Services.Entities;

namespace AvaloniaApplication7.Services.Contracts;

public interface IDbService
{
    IEnumerable<Singer> GetAllSingers();
    IEnumerable<Song> GetSingerSongs(int songId);
    void Init();
}